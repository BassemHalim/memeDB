"use client";
import Timeline from "@/components/ui/Timeline";
import { useMemes } from "@/hooks/useMemes";
import { useRouter } from "@/i18n/navigation";
import { Meme } from "@/types/Meme";
import { sendEvent } from "@/utils/googleAnalytics";
import { ChevronLeft, ChevronRight, Search as SearchIcon } from "lucide-react";
import { useTranslations } from "next-intl";
import { FormEvent, useEffect, useRef, useState } from "react";
import { cn } from "../utils/tailwind";
import { Button } from "./ui/button";
import { Toggle } from "./ui/toggle";

interface Tag {
    name: string;
    pressed: boolean;
}

export default function SearchComponent({
    query,
    selectedTags,
}: {
    query?: string;
    selectedTags: string[];
}) {
    const [memes, setMemes] = useState<Meme[]>([]);
    const [tags, setTags] = useState<Tag[]>(
        [...new Set(memes.map((meme) => meme.tags).flat()).values()]
            .sort((e1, e2) => {
                if (e1 == e2) return 0;
                else if (e1 < e2) return 1;
                return -1;
            })
            .map((tag) => {
                if (selectedTags.includes(tag))
                    return { name: tag, pressed: true };
                return { name: tag, pressed: false };
            })
    );
    const [error, setError] = useState(false);
    const [noMatch, setNoMatch] = useState(false);
    const [isLoading, setIsLoading] = useState<boolean>(false);
    const router = useRouter();
    const tagsRef = useRef<HTMLDivElement>(null);
    const [tagsOverflow, setTagsOverflow] = useState(true);
    const { memes: recommendedMemes } = useMemes(undefined, 10);
    const t = useTranslations("search");

    useEffect(() => {
        if (!query) {
            return;
        }
        sendEvent("memes_search", { query: query });
    }, []);
    useEffect(() => {
        const checkOverflow = () => {
            if (tagsRef.current) {
                setTagsOverflow(tagsRef.current.scrollWidth > screen.width);
            } else {
                setTagsOverflow(true);
            }
        };

        checkOverflow();

        // Add resize listener to recheck when window size changes
        window.addEventListener("resize", checkOverflow);
        return () => window.removeEventListener("resize", checkOverflow);
    }, [tags, memes]);

    function search(query: string, tags?: string[]) {
        // go to /search?query={query}&tags={tags}
        type searchQuery = {
            query: string;
            tags?: string;
        };
        const queryObj: searchQuery = {
            query: query,
        };
        if (tags && tags.length > 0) {
            queryObj.tags = tags.join(",");
        }
        router.push({ pathname: "/search", query: queryObj });
    }

    function onSubmit(e: FormEvent) {
        e.preventDefault();
        const form = e.target as HTMLFormElement;
        const data = new FormData(form);
        const query = data.get("query")?.toString();
        if (!query || query.length == 0) {
            return;
        }
        // update page search params
        search(query);
    }
    async function searchMemes(query: string | null) {
        setNoMatch(false);
        setError(false);
        if (!query || query.length == 0) {
            return;
        }
        setIsLoading(true);
        // updated searchBar value
        const input = document.getElementById("searchBar") as HTMLInputElement;
        input.value = query;

        // make fetch request to /api/memes?query
        const url = new URL(
            "/api/memes/search",
            process.env.NEXT_PUBLIC_API_HOST
        );
        url.searchParams.append("query", query);
        url.searchParams.append("pageSize", "100");
        try {
            const resp = await fetch(url);
            if (!resp.ok) {
                throw new Error("Failed to fetch memes");
            }
            const memesResp = await resp.json();
            if (memesResp.memes) {
                const newMemes: Meme[] = memesResp.memes;
                setMemes(newMemes);
            } else {
                setNoMatch(true);
                setIsLoading(false);
                return;
            }
            // form.reset();
            setError(false);
            setNoMatch(false);
            setIsLoading(false);
        } catch (err) {
            setError(true);
            setIsLoading(false);
            console.log(err);
        }
    }

    useEffect(() => {
        setTags(
            [...new Set(memes.map((meme) => meme.tags).flat()).values()].map(
                (tag) => {
                    return { name: tag, pressed: selectedTags.includes(tag) };
                }
            )
        );
        setTimeout(() => {
            if (tagsRef.current) {
                setTagsOverflow(tagsRef.current.scrollWidth > screen.width);
            }
        }, 5);
    }, [memes, selectedTags]);

    useEffect(() => {
        if (!query) {
            return;
        }
        const searchQuery = [query, ...selectedTags].join(" ");
        searchMemes(searchQuery);
    }, [query, selectedTags]);

    return (
        <section className="w-full flex-1 mt-4 gap-2 flex flex-col overflow-hidden">
            <div className="w-full max-w-2xl relative text-gray-800 mx-auto px-2">
                <form onSubmit={onSubmit}>
                    <input
                        required
                        name="query"
                        id="searchBar"
                        type="text"
                        placeholder={t("meme-search")}
                        className={`w-full px-6 py-4 rounded-lg text-lg shadow-lg pr-14`}
                    />
                    <button type="submit">
                        <SearchIcon className="absolute right-4 top-0 mt-4 text-gray-400 h-6 w-6" />
                    </button>
                </form>
            </div>
            {noMatch ? (
                <div className="h-12 flex flex-col justify-center">
                    <h2 className="text-xl font-bold text-center p-4">
                        {t("no-matches")}
                    </h2>
                </div>
            ) : null}
            {error ? (
                <div className="h-12 flex flex-col justify-center mt-8">
                    <h2 className="text-xl font-bold text-center p-4">
                        We encountered a problem while looking for your memes
                    </h2>
                </div>
            ) : null}
            <div className="font-bold relative mx-2 h-10 flex items-center justify-center group">
                <div
                    className={cn(
                        "flex  gap-2 justify-center",
                        tagsOverflow ? "absolute" : ""
                    )}
                    ref={tagsRef}
                    style={{ left: 0 }}
                >
                    {tags.map((tag) => (
                        <Toggle
                            pressed={tag.pressed}
                            onPressedChange={(pressed) => {
                                const newTags = tags.map((t) => {
                                    if (t.name === tag.name) {
                                        return {
                                            name: t.name,
                                            pressed: pressed,
                                        };
                                    }
                                    return t;
                                });
                                setTags(newTags);
                                const pressedTags = newTags
                                    .filter((tag) => tag.pressed)
                                    .map((tag) => tag.name);
                                search(query!, pressedTags);
                            }}
                            key={tag.name}
                            className="p-2 text-center text-nowrap bg-primary text-secondary"
                        >
                            {tag.name}
                        </Toggle>
                    ))}
                </div>
                {tagsOverflow && (
                    <>
                        <Button
                            className="md:hidden group-hover:inline-block absolute right-4 top-1 rounded-full p-1 h-8 w-8 bg-secondary text-primary hover:bg-secondary/80"
                            onClick={() => {
                                const carousel = tagsRef.current;
                                if (carousel) {
                                    const incr = screen.width / 2;
                                    const carouselWidth = parseInt(
                                        getComputedStyle(carousel).width,
                                        10
                                    );
                                    carousel.style.left = `${Math.max(
                                        parseInt(carousel.style.left, 10) -
                                            incr,
                                        -carouselWidth + screen.width - 140
                                    )}px`;
                                }
                            }}
                        >
                            <ChevronRight className="h-4 w-4 m-auto" />
                        </Button>
                        <Button
                            className="md:hidden group-hover:inline-block absolute left-4 top-1 rounded-full p-1 h-8 w-8 bg-secondary text-primary hover:bg-secondary/80"
                            onClick={() => {
                                const carousel = tagsRef.current;
                                if (carousel) {
                                    const incr = screen.width / 2;
                                    carousel.style.left = `${Math.min(
                                        parseInt(carousel.style.left, 10) +
                                            incr,
                                        70
                                    )}px`;
                                }
                            }}
                        >
                            <ChevronLeft className="h-4 w-4 m-auto" />
                        </Button>
                    </>
                )}
            </div>
            {noMatch ? (
                <>
                    <p className="text-center">{t("you-might-like")}</p>
                    <Timeline
                        memes={recommendedMemes}
                        isLoading={false}
                        hasMore={false}
                    />
                </>
            ) : (
                <Timeline memes={memes} isLoading={isLoading} />
            )}
        </section>
    );
}
