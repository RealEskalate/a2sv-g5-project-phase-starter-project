import React from "react";

interface RectangleProps {
  backgroundColor: string;
  src: string;
}

export default function Rectangle({ backgroundColor, src }: RectangleProps) {
  return (
    <div
      className={`w-14 h-14 bg-[rgba(${backgroundColor})] rounded-3xl flex justify-center items-center`}
    >
      <img src={src} alt="" />
    </div>
  );
}
