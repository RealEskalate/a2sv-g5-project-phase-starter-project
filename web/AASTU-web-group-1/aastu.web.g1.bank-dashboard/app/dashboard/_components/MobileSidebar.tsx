import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetTrigger,
} from "@/components/ui/sheet";
import { sidebarLinks } from "@/constants";
import { useUser } from "@/contexts/UserContext";
import { cn } from "@/lib/utils";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

const MobileSidebar = () => {
  const pathname = usePathname();
  const { isDarkMode } = useUser();

  return (
    <section
      className={cn(
        "w-full max-w-[264px]",
        isDarkMode ? "bg-gray-800" : "bg-white"
      )}
    >
      <Sheet>
        <SheetTrigger>
          <Image
            src="/icons/hamburger.svg"
            width={30}
            height={30}
            alt="Hamburger"
            className="cursor-pointer"
          />
        </SheetTrigger>
        <SheetContent
          side="left"
          className={cn(
            "border-none",
            isDarkMode ? "bg-gray-900 text-white" : "bg-white text-black"
          )}
        >
          <Link
            href="/"
            className="cursor-pointer flex items-center gap-1 px-4"
          >
            <Image src="/icons/logo.png" width={20} height={20} alt="logo" />
            <h1
              className={cn(
                "font-[900] text-[1.4rem]",
                isDarkMode ? "text-white" : "text-primaryBlack"
              )}
            >
              BankDash.
            </h1>
          </Link>
          <div>
            <SheetClose asChild>
              <nav className="flex h-full flex-col gap-6 pt-16">
                {sidebarLinks.map((link) => {
                  const isActive =
                    pathname === link.route ||
                    pathname.startsWith(`/dashboard${link.route}/`);

                  return (
                    <SheetClose asChild key={link.route}>
                      <Link
                        href={link.route}
                        key={link.title}
                        className={cn(
                          "flex gap-6 items-center py-1 md:p-3 2xl:px-4 pl-0 justify-start xl:justify-start"
                        )}
                      >
                        <Image
                          src={link.icon}
                          alt={link.title}
                          width={20}
                          height={20}
                          className={cn({
                            "filter-custom-blue": isActive,
                            invert: isDarkMode, // Invert the icon color in dark mode
                          })}
                        />
                        <p
                          className={cn("text-sm font-semibold", {
                            "text-primaryBlue": isActive && !isDarkMode,
                            "text-blue-500": isActive && isDarkMode,
                            "text-[#B1B1B1]": !isActive && !isDarkMode,
                            "text-gray-400": !isActive && isDarkMode,
                          })}
                        >
                          {link.title}
                        </p>
                      </Link>
                    </SheetClose>
                  );
                })}
              </nav>
            </SheetClose>
          </div>
        </SheetContent>
      </Sheet>
    </section>
  );
};

export default MobileSidebar;
