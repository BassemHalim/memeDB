import Footer from "@/components/Footer";
import Header from "@/components/Header";
import "@ant-design/v5-patch-for-react-19";
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
    title: "Qasr El Memez",
    description: "The home of your memes",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body className={` antialiased flex flex-col min-h-screen`}>
                <Header />
                <main className="grow flex flex-col items-center justify-start w-full">
                    {children}
                </main>
                <Footer />
            </body>
        </html>
    );
}
