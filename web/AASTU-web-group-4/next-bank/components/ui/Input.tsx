import React from 'react';

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
}

const Input: React.FC<InputProps> = ({ label, ...props }) => (
  <div className="flex flex-col">
    <label className="text-sm font-medium">{label}</label>
    <input
      {...props}
      className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
    />
  </div>
);

export default Input;
