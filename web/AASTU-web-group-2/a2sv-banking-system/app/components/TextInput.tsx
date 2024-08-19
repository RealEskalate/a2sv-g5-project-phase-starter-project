import React from 'react';

interface TextInputProps {
  id: string;
  label: string;
  value: string;
  readOnly?: boolean;
}

const TextInput: React.FC<TextInputProps> = ({ id, label, value, readOnly }) => {
  return (
    <div>
      <label htmlFor={id} className="block text-sm font-medium text-[#232323]">{label}</label>
      <input
        type="text"
        id={id}
        title={label}
        className="mt-1 block w-full border border-[#DFEAF2] rounded-full shadow-sm px-4 py-2 text-[#718EBF]"
        value={value}
        readOnly={readOnly}
      />
    </div>
  );
};

export default TextInput;
