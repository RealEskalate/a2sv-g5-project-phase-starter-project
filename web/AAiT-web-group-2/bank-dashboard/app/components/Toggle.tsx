import { useState } from "react";

const ToggleButton = ({enabled, onEnable}: {enabled: boolean, onEnable: () => void}) => {

  return (
    <button
      onClick={() => onEnable()}
      className={`relative inline-flex items-center h-6 rounded-full w-11 transition-colors duration-300 focus:outline-none ${
        enabled ? "bg-custom-greenish" : "bg-gray-200"
      }`}
    >
      <span
        className={`inline-block w-[22px] h-[22px] transform bg-white rounded-full transition-transform duration-300 ${
          enabled ? "translate-x-5" : "translate-x"
        }`}
      />
    </button>
  );
};

export default ToggleButton;
