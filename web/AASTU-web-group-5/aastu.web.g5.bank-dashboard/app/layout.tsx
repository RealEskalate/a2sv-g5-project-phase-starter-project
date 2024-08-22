'use client';
<<<<<<< Updated upstream
import { useState, useEffect } from "react";
import { Inter } from "next/font/google";
import { SessionProvider } from "next-auth/react";
import { Provider } from 'react-redux'; 
import store from './redux/store'; 
import "./globals.css";
import NavBar from "./components/common/navBar";
import Sidebar from "./components/common/sideBar";
=======

import { useState, useEffect } from "react";
import { Inter } from "next/font/google";
import { SessionProvider } from "next-auth/react";
import { Provider } from 'react-redux';
// import store from './redux/store';
import "./globals.css";
import NavBar from "./components/common/navBar";
import SideBar from "./components/common/sideBar";
import { useSession } from "next-auth/react";
>>>>>>> Stashed changes

const inter = Inter({ subsets: ["latin"] });

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
<<<<<<< Updated upstream
	const [isSidebarVisible, setIsSidebarVisible] = useState(false);
	const [darkmode, setDarkmode] = useState(false); // Add state for dark mode
=======
  const [isSidebarVisible, setIsSidebarVisible] = useState(false);
  const [darkmode, setDarkmode] = useState(false);

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', darkmode ? 'dark' : 'light');
  }, [darkmode]);
>>>>>>> Stashed changes

	const toggleSidebar = () => {
		setIsSidebarVisible(!isSidebarVisible);
	};

<<<<<<< Updated upstream
	const toggleDarkMode = () => {
		setDarkmode(!darkmode);
	};

	useEffect(() => {
		if (darkmode) {
			document.documentElement.setAttribute('data-theme', 'dark');
		} else {
			document.documentElement.removeAttribute('data-theme');
		}
	}, [darkmode]);

	return (
		<html lang="en">
			<body className={inter.className}>
				<SessionProvider>
					<Provider store={store}>
						<div className="min-h-screen bg-slate-200 sm:grid sm:grid-cols-[200px_1fr] md:grid-cols-[250px_1fr]">
							<div className={`fixed inset-0 bg-white z-50 sm:static sm:block ${isSidebarVisible ? 'block' : 'hidden'}`}>
								<Sidebar isSidebarVisible={isSidebarVisible} toggleSidebar={toggleSidebar} />
							</div>
							<div className="flex flex-col w-full">
								<NavBar toggleSidebar={toggleSidebar} isSidebarVisible={isSidebarVisible} toggleDarkMode={toggleDarkMode} darkmode={darkmode} />
								<main>{children}</main>
							</div>
						</div>
					</Provider>
				</SessionProvider>
			</body>
		</html>
	);
=======
  const toggleDarkMode = () => {
    setDarkmode(!darkmode);
  };

  return (
    <html lang="en">
      <body className={inter.className}>
        <SessionProvider>
          {/* <Provider store={store}> */}
            <SessionWrapper>
              <div className={`min-h-screen flex ${darkmode ? 'dark' : ''}`}>
                <SidebarWrapper
                  isSidebarVisible={isSidebarVisible}
                  toggleSidebar={toggleSidebar}
                />
                <div className="flex flex-col flex-1 transition-all duration-300">
                  <NavBar
                    toggleSidebar={toggleSidebar}
                    isSidebarVisible={isSidebarVisible}
                    // toggleDarkMode={toggleDarkMode}
                    // darkmode={darkmode}
                  />
                  <main >
                    {children}
                  </main>
                </div>
              </div>
            </SessionWrapper>
          {/* </Provider> */}
        </SessionProvider>
      </body>
    </html>
  );
>>>>>>> Stashed changes
}

function SessionWrapper({ children }: { children: React.ReactNode }) {
  const { status } = useSession();

  if (status === "loading") {
    return <div>Loading...</div>;
  }

  return <>{children}</>;
}

function SidebarWrapper({ isSidebarVisible, toggleSidebar }: { isSidebarVisible: boolean, toggleSidebar: () => void }) {
  const { status } = useSession();

  return (
    status === "authenticated" && (
      <div className={`fixed inset-0 bg-white z-50 sm:static sm:block ${isSidebarVisible ? 'block' : 'hidden'}`}>
        <SideBar
          isSidebarVisible={isSidebarVisible}
          toggleSidebar={toggleSidebar}
        />
      </div>
    )
  );
}
// 