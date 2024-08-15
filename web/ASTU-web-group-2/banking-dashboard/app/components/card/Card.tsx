import React from "react";

const Card = ({
  title,
  className,
  children,

}: {
  title: string;
  className?: string;
  children: React.ReactNode;
}) => {
  return (
    <div className={`${className}`}>
      <div>
        <p className="font-semibold my-3 text-[18px] md:text-[18px] lg:text[22px] text-[#343C6A]">
          {title}
        </p>
      </div>
      {children}
    </div>
  );
};

export default Card;
