// components/TextInput.tsx
import React from 'react';

interface TextInputProps {
  id: string;
  label: string;
  value: string;
  placeholder?: string;
  readOnly?: boolean; 
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

const TextInput: React.FC<TextInputProps> = ({ id, label, value, placeholder, readOnly = false, onChange }) => {
  return (
    <div>
      <label htmlFor={id} className="block text-sm font-medium text-[#232323]">{label}</label>
      <input
        type="text"
        id={id}
        title={label}
        className="mt-1 block w-full border border-[#DFEAF2] rounded-full shadow-sm px-4 py-2 text-[#718EBF]"
        value={value}
        placeholder={placeholder}
        readOnly={readOnly}
        onChange={onChange}
      />
    </div>
  );
};

export default TextInput;
