"use client";

import MemeEditor from "@/components/ui/memeEditor/memeEditor";
import { useTranslations } from "next-intl";

export default function Page() {
    const t = useTranslations("memeGenerator");
    return (
        <main>
            <div className="max-w-sm md:max-w-md lg:max-w-lg mx-auto mb-4">
                <h1 className="font-bold text-2xl sm:text-3xl md:text-4xl text-center p-4 text-wrap">
                    {t("title")}
                </h1>
            </div>
            <MemeEditor />
            <div className="p-4 max-w-[90vw] flex flex-col justify-center items-center mx-auto">
                <p>{t("subtitle")}</p>
                <h2 className="font-semibold text-2xl p-2">
                    {t("instructions-title")}
                </h2>
                <ol className="list-decimal">
                    <li>
                        {t.rich("step-1", {
                            strong: (chunks) => <strong>{chunks}</strong>,
                        })}
                    </li>
                    <li>
                        {t.rich("step-2", {
                            strong: (chunks) => <strong>{chunks}</strong>,
                        })}
                    </li>
                    <li>
                        {t.rich("step-3", {
                            strong: (chunks) => <strong>{chunks}</strong>,
                        })}
                    </li>
                </ol>
            </div>
        </main>
    );
}
