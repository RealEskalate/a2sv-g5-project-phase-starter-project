'use client';
import { useState, useEffect } from "react";
import { Inter } from "next/font/google";
import { SessionProvider } from "next-auth/react";
import { Provider } from 'react-redux'; 
import store from './redux/store'; 
import "./globals.css";
import NavBar from "./components/common/navBar";
import Sidebar from "./components/common/sideBar";

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	const [isSidebarVisible, setIsSidebarVisible] = useState(false);
	const [darkmode, setDarkmode] = useState(false); // Add state for dark mode
	const isAuthPage = window.location.pathname.startsWith("/auth");
	useEffect(() => {
		if (darkmode) {
			document.documentElement.setAttribute('data-theme', 'dark');
		} else {
			document.documentElement.removeAttribute('data-theme');
		}
	}, [darkmode]);


	const toggleSidebar = () => {
		setIsSidebarVisible(!isSidebarVisible);
	};

	const toggleDarkMode = () => {
		setDarkmode(!darkmode);
	};




	return (
		<html lang="en">
			<body className={inter.className}>
				<SessionProvider>
					<Provider store={store}>
						{!isAuthPage && (
							<div className="min-h-screen bg-slate-200 sm:grid sm:grid-cols-[200px_1fr] md:grid-cols-[250px_1fr]">
								<div className={`fixed inset-0 bg-white z-50 sm:static sm:block ${isSidebarVisible ? 'block' : 'hidden'}`}>
									<Sidebar isSidebarVisible={isSidebarVisible} toggleSidebar={toggleSidebar} />
								</div>
								<div className="flex flex-col w-full">
									<NavBar toggleSidebar={toggleSidebar} isSidebarVisible={isSidebarVisible} />
									<main>{children}</main>
								</div>
							</div>
						)}
						{isAuthPage && (
							<main>{children}</main> 
						)}
					</Provider>
				</SessionProvider>
			</body>
		</html>
	);
}
