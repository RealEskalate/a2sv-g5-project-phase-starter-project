import Header from "@/components/dashboard-layout/header";
import SideBar from "@/components/dashboard-layout/sidebar";
import Image from "next/image";
import React from "react";

type Props = {
    children: React.ReactNode;
};

export default function DashboardLayout({ children }: Props) {
    return (
        <div className="wrapper">
            <div className="sidebar hidden md:block fixed ">
                <div className="flex gap-x-2 h-[100px] items-center pl-10">
                    <Image src='/icons/logo.svg' alt="Logo" width={36} height={36} />
                    <h1 className="font-[900] text-2xl">BankDash.</h1>
                </div>
                <SideBar />
            </div>
            <div className="RightSide sm:ml-[250px]">
                <div className="topheader">
                    <Header />
                </div>
                <div className="content">
                    <main className="bg-[#E6EFF5] min-h-screen">{children}</main>
                </div>
            </div>
        </div>
    );
}
