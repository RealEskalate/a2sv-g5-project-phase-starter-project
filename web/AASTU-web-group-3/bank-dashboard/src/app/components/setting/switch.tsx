import React, { useState } from 'react';

interface SwitchButtonProps {
  isOn?: boolean;
  onToggle?: (state: boolean) => void;
}

const SwitchButton: React.FC<SwitchButtonProps> = ({ isOn = false, onToggle }) => {
  const [toggled, setToggled] = useState(isOn);

  const handleToggle = () => {
    console.log("toggled",!toggled)
    const newState = !toggled;
    setToggled(newState);
    if (onToggle) {
      onToggle(newState);
    }
  };

  return (
    <div onClick={handleToggle}
      className={`flex items-center cursor-pointer w-14  h-8 p-1 rounded-full ${toggled ? 'bg-[#16DBCC]' : 'bg-gray-300'}`}
    >
      <div
        className={`bg-white w-6 h-6 rounded-full shadow-md transform ${toggled ? 'translate-x-6' : ''} transition-transform duration-300 ease-in-out`}
      >
      </div>
    </div>
  );
};

export default SwitchButton;
