"use client";
import { useSession } from "next-auth/react";
import Investments from "./Investments/page";
import Signin from "./auth/signin/page";
import { useSelector } from "react-redux";
import { RootState } from "@/app/redux/store";

export default function Home() {
	const { status } = useSession();
	const darkMode = useSelector((state: RootState) => state.theme.darkMode);

    return (
        <div className={`p-4 ${darkMode ? 'bg-gray-900 text-white' : 'bg-white text-neutral-800'}`}>
            {/* Render Investments if authenticated, otherwise render Signin */}
            {status === "authenticated" ? <Investments /> : <Signin />}
        </div>
    );
}
