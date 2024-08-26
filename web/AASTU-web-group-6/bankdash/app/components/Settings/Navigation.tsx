import React from "react";

interface MenuItem {
  label: string;
  section: string;
}

interface NavigationProps {
  activeSection: string;
  setActiveSection: (section: string) => void;
}

const Navigation: React.FC<NavigationProps> = ({
  activeSection,
  setActiveSection,
}) => {
  const menuItems: MenuItem[] = [
    { label: "Edit Profile", section: "editprofile" },
    { label: "Preferences", section: "preference" },
    { label: "Security", section: "security" },
  ];

  return (
    <div className="w-full max-h-16 flex font-Inter justify-between items-center border-b dark:border-gray-500 xs:mt-6 sm:mt-2 pt-4 px-2 min-h-6">
      <div className="flex xxs:px-1 xxs:w-full xxs:justify-between xxs:gap-2 xs:justify-normal xs:px-4 md:px-7 md:gap-8">
        {menuItems.map((item, index) => (
          <button
            key={index}
            onClick={() => setActiveSection(item.section)}
            className={`border-b-[6px] rounded cursor-pointer font-Inter xxs:w-fit xxs:text-base md:text-lg md:w-24 dark:text-white ${
              activeSection === item.section
                ? "border-[#1814F3]"
                : "border-transparent"
            }`}
          >
            {item.label}
          </button>
        ))}
      </div>
    </div>
  );
};

export default Navigation;
