"use client";
import Timeline from "@/components/ui/Timeline";
import { Meme } from "@/types/Meme";
import { Search as SearchIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { FormEvent, useEffect, useState } from "react";

export default function SearchComponent({ query }: { query: string }) {
    const [memes, setMemes] = useState<Meme[]>([]);
    const [error, setError] = useState(false);
    const [noMatch, setNoMatch] = useState(false);
    const [isLoading, setIsLoading] = useState<boolean>(false);
    const router = useRouter();

    function onSubmit(e: FormEvent) {
        e.preventDefault();
        const form = e.target as HTMLFormElement;
        const data = new FormData(form);
        const query = data.get("query")?.toString();
        if (!query || query.length == 0) {
            return;
        }

        // update page search params
        // go to /search?query={query}
        const url = new URL("search", window.location.origin);
        url.searchParams.append("query", query);
        router.push(url.href);
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
        searchMemes(query);
    }, [query]);
    return (
        <section className="w-full flex-1 mt-4">
            <div className="w-full max-w-2xl relative text-gray-800 mx-auto px-2">
                <form onSubmit={onSubmit}>
                    <input
                        required
                        name="query"
                        id="searchBar"
                        type="text"
                        placeholder="Search memes..."
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
                        Sorry, we couldn&apos;t find relevant Memes for this
                        search
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
            <Timeline memes={memes} isLoading={isLoading} />
        </section>
    );
}
