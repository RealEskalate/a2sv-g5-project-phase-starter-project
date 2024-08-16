import React from 'react';

interface NotificationToggleProps {
  id: string;
  label: string;
  checked: boolean;
  onChange: (checked: boolean) => void;
}

const NotificationToggle: React.FC<NotificationToggleProps> = ({ id, label, checked, onChange }) => {
  return (
    <div className="flex items-center gap-4 mb-4">
      <div className={`relative w-12 h-6 ${checked ? 'bg-[#16DBCC]' : 'bg-[#DFEAF2]'} rounded-full flex items-center`}>
        <div
          className={`absolute h-6 w-6 bg-white rounded-full ${checked ? 'right-0.5' : 'left-0.5'}`}
        ></div>
      </div>
      <label htmlFor={id} className="text-gray-700">{label}</label>
      <input
        type="checkbox"
        id={id}
        className="hidden"
        checked={checked}
        onChange={(e) => onChange(e.target.checked)}
      />
    </div>
  );
};

export default NotificationToggle;
