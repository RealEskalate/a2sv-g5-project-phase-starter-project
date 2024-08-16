import React from 'react';

interface ToggleSwitchProps {
  id: string;
  label: string;
  checked: boolean;
  onChange: (checked: boolean) => void;
}

const ToggleSwitch: React.FC<ToggleSwitchProps> = ({ id, label, checked, onChange }) => {
  return (
    <label htmlFor={id} className="inline-flex items-center cursor-pointer gap-[1rem]">
      <input
        type="checkbox"
        id={id}
        className="sr-only peer"
        checked={checked}
        onChange={() => onChange(!checked)}
      />
      <div
        className={`relative max-md:w-11 max-md:h-6 w-[56px] h-[30.71px] rounded-full transition-colors duration-200 ease-in-out ${
          checked ? 'bg-[#16DBCC]' : 'bg-gray-200'
        }`}
      >
        <div
          className={`absolute left-0 top-0.5 max-md:w-5 max-md:h-5 w-[27px] h-[27px] bg-white rounded-full transition-transform duration-200 ease-in-out ${
            checked ? 'max-md:translate-x-5 translate-x-7 ' : 'max-md:translate-x-0 translate-x-1'
          }`}
        ></div>
      </div>
      <span className="font-[400] text-[16px] text-[#232323]">{label}</span>
    </label>
  );
};

export default ToggleSwitch;
