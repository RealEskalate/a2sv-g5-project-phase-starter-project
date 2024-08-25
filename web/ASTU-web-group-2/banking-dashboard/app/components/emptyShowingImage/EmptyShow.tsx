import React from "react";
import Image from "next/image";
interface EmptyMessageProps {
  text: string;
}

const EmptyShow: React.FC<EmptyMessageProps> = ({ text }) => {
  return (
    <div className="flex flex-col items-center justify-center min-h-[200px] p-6 rounded-lg w-[100%]">
      <Image
        src="/assets/emptyimage/emptyimage.png"
        width={200}
        height={100}
        alt="empty showing image"
      />
      <p className="text-gray text-center font-semibold">{text}</p>
    </div>
  );
};

export default EmptyShow;
