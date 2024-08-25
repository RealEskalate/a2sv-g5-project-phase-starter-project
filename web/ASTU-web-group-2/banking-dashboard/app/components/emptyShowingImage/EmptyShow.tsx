import React from "react";
import Image from "next/image";
interface ImageWithTextProps {
  text: string;
}

const EmptyShow: React.FC<ImageWithTextProps> = ({ text }) => {
  return (
    <div className="text-center flex justify-center flex-col">
      <Image
        src="/assets/emptyImage/empty.png"
        width={200}
        height={100}
        alt="empty showing image"
      />
      <p className="mt-2 text-lg font-semibold">{text}</p>
    </div>
  );
};

export default EmptyShow;
