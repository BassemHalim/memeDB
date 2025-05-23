import { getTranslations } from "next-intl/server";

export async function generateMetadata({
    params,
}: {
    params: Promise<{ locale: string }>;
}) {
    const { locale } = await params;
    const t = await getTranslations({
        locale: locale,
        namespace: "generatorMetadata",
    });

    return {
        title: t("title"),
        description: t("description"),
        openGraph: {
            type: "website",
            url: `https://qasrelmemez.com/${locale}/generator`,
            title: t("title"),
            description: t("description"),
            images: [
                {
                    url: "https://qasrelmemez.com/og.png",
                    width: 1200,
                    height: 630,
                    alt: "Qasr El Memez",
                },
            ],
        },
        alternates: {
            canonical: `https://qasrelmemez.com/generator`,
        },
    };
}

export default function GeneratorLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return <section>{children}</section>;
}
