"use client";

import { useEffect, useState } from "react";
import { Inter } from "next/font/google";
import { SessionProvider } from "next-auth/react";
import { Provider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import { store, persistor } from "./redux/store";
import "./globals.css";
import NavBar from "./components/common/navBar";
import SideBar from "./components/common/sideBar";
import { useSession } from "next-auth/react";
import { useSelector } from "react-redux";
import { RootState } from "./redux/store";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html lang="en">
			<body className={`${inter.className} dark:bg-gray-900`}>
				<SessionProvider>
					<Provider store={store}>
						<PersistGate loading={null} persistor={persistor}>
							<LayoutContent>{children}</LayoutContent>
						</PersistGate>
					</Provider>
				</SessionProvider>
			</body>
		</html>
	);
}

function LayoutContent({ children }: { children: React.ReactNode }) {
	const { status } = useSession();
	const darkmode = useSelector((state: RootState) => state.theme.darkMode);
	const [isSidebarVisible, setIsSidebarVisible] = useState(false);

	useEffect(() => {
		document.documentElement.setAttribute(
			"data-theme",
			darkmode ? "dark" : "light"
		);
	}, [darkmode]);

	const toggleSidebar = () => {
		setIsSidebarVisible(!isSidebarVisible);
	};

	const darkClass = darkmode ? "dark" : "";
	return (
		<div className={`min-h-screen flex ${darkClass}`}>
			{status === "authenticated" && (
				<SidebarWrapper
					isSidebarVisible={isSidebarVisible}
					toggleSidebar={toggleSidebar}
				/>
			)}
			<div className="flex flex-col flex-1 transition-all duration-300">
				<NavBar
					toggleSidebar={toggleSidebar}
					isSidebarVisible={isSidebarVisible}
				/>
				<main>{children}</main>
			</div>
		</div>
	);
}

function SidebarWrapper({
	isSidebarVisible,
	toggleSidebar,
}: {
	isSidebarVisible: boolean;
	toggleSidebar: () => void;
}) {
	const { status } = useSession();

	return (
		status === "authenticated" && (
			<div
				className={`fixed inset-0 bg-white z-50 sm:static sm:block ${
					isSidebarVisible ? "block" : "hidden"
				} dark:bg-gray-800 dark:text-white pr-10`}
			>
				<SideBar
					isSidebarVisible={isSidebarVisible}
					toggleSidebar={toggleSidebar}
				/>
			</div>
		)
	);
}
