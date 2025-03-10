import { getUserLocale, setUserLocale } from "@/services/locale";
import { Languages } from "lucide-react";
import { useEffect, useState } from "react";
import { Button } from "./button";

export default function LanguageSwitch() {
    const [arabic, setArabic] = useState(false);
    useEffect(() => {
        getUserLocale().then((val) => {
            if (val == "ar") {
                setArabic(true);
            } else {
                setArabic(false);
            }
        });
    }, []);

    // const label = arabic ? "EN" : "AR";
    return (
        <Button
            aria-label="language button"
            className="bg-primary text-secondary "
            onClick={() => {
                setArabic(!arabic);
                setUserLocale(!arabic ? "ar" : "en");
            }}
        >
            <Languages />
        </Button>
    );
}
