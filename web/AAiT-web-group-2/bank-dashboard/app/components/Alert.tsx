
import React, { useState } from 'react';

interface AlertProps {
  type: 'success' | 'error';
  message: string;
  duration: number;
  onClose?: () => void
}

const Alert: React.FC<AlertProps> = ({ message, type, duration = 5000, onClose }) => {
    const [visible, setVisible] = useState(true);
    const alertStyles = {
        success: 'bg-green-100 text-green-800 border-green-400',
        error: 'bg-red-100 text-red-800 border-red-400',
      };
  
    React.useEffect(() => {
      if (duration && visible) {
        const timer = setTimeout(() => {
          setVisible(false);
          if (onClose) onClose();
        }, duration);
  
        return () => clearTimeout(timer); 
      }
    }, [duration, visible, onClose]);
  
    if (!visible) return null;
  return (
    <div
      className={`p-4 mb-4 z-20 border absolute bottom-10 left-10 shadow-lg rounded-lg ${alertStyles[type]}`}
      role="alert"
    >
      <strong className="font-medium">{type === 'success' ? 'Success!' : 'Error!'}</strong>
      <p className="mt-2 text-wrap">{message}</p>
    </div>
  );
};

export default Alert;

