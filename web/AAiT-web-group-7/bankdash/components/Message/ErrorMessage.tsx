import React from 'react';

interface MessageError {
  message: string | undefined;
}

const ErrorMessage: React.FC<MessageError> = ({ message }) => {
  return (
    <p className="text-red-600 flex text-xs font-semibold gap-1">
      {message && (
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth="1.5"
          stroke="currentColor"
          className="size-4"
        >
          <path
            strokeLinecap="round"
            stroke-linejoin="round"
            d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
          />
        </svg>
      )}
      {message}
    </p>
  );
};

export default ErrorMessage;