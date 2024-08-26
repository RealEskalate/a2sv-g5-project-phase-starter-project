"use client";
import React, { useState } from "react";

interface NotificationToggleProps {
  id: string;
  label: string;
  checked: boolean;
  onChange: (checked: boolean) => void;
}

const NotificationToggle: React.FC<NotificationToggleProps> = ({ id, label, checked, onChange }) => {
  const [enabled, setEnabled] = useState(checked);

  const handleToggle = () => {
    const newChecked = !enabled;
    setEnabled(newChecked);
    onChange(newChecked);
  };

  return (
   <div className="flex items-center gap-3">
  <input
    type="checkbox"
    id={id}
    className="peer hidden"
    checked={enabled}
    onChange={handleToggle}
  />
  <label
    htmlFor={id}
    className={`cursor-pointer rounded-full w-12 h-6 flex items-center relative transition-colors duration-300 shadow-md ${
      enabled ? "bg-[#16dbcc] dark:bg-[#1814f3]" : "bg-gray-200 dark:bg-gray-600"
    }`}
  >
    <span
      className={`bg-white w-5 h-5 rounded-full transition-transform duration-300 transform shadow-sm ${
        enabled ? "translate-x-6" : ""
      } dark:bg-gray-300`}
    ></span>
  </label>

  <label htmlFor={id} className="text-xs text-[#232323] lg:text-base dark:text-[#9faaeb]">
    {label}
  </label>
</div>
  );
};

export default NotificationToggle;
