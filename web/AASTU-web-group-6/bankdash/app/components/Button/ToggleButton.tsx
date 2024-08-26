import React, { useState } from "react";

type ToggleButtonValue = {
  onToggle: (checked: boolean) => void;
  initialChecked?: boolean;
};

const ToggleButton: React.FC<ToggleButtonValue> = ({
  onToggle,
  initialChecked = false,
}) => {
  const [checked, setChecked] = useState(initialChecked);

  const handleToggle = () => {
    setChecked(!checked);
    onToggle(!checked);
  };

  return (
    <div className="relative inline-block w-12 min-w-12 h-6 px-[1px]">
      <input
        type="checkbox"
        checked={checked}
        onChange={handleToggle}
        className="opacity-0 w-0 h-0"
      />
      <span
        className={`absolute cursor-pointer top-0 left-0 right-0 bottom-0 bg-gray-300 dark:bg-gray-600 rounded-full transition-all duration-300 ${
          checked
            ? "bg-teal-400 dark:bg-teal-500"
            : "bg-gray-300 dark:bg-gray-600"
        }`}
      >
        <span
          className={`absolute left-1 top-1 bg-white dark:bg-gray-200 w-4 h-4 rounded-full transition-transform duration-300 transform ${
            checked ? "translate-x-6" : "translate-x-0"
          }`}
        ></span>
      </span>
    </div>
  );
};

export default ToggleButton;
