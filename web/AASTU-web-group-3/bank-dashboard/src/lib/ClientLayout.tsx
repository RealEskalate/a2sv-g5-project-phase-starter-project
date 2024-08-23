import { usePathname } from "next/navigation";
import Navbar from "@/app/components/NavBar";
import Sidebar from "@/app/components/SideBar";
import { ReactNode } from "react";

export default function ClientLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname();

  const showSidebarAndNavbar = pathname !== "/auth/signin" && pathname !== "/auth/signup";
  return (
    <>
      {showSidebarAndNavbar && <Sidebar />}
      <div className="flex flex-col flex-grow h-full overflow-hidden md:w-4/5 lg:w-4/5">
        {showSidebarAndNavbar && <Navbar />}
        <main className="flex-grow overflow-y-auto bg-[#F5F7FA] p-1">
          {children}
        </main>
      </div>
    </>
  );
}
