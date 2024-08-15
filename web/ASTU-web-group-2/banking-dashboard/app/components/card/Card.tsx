import React from "react";

const Card = ({
  title,
  children,
}: {
  title: string;
  children: React.ReactNode;
}) => {
  return (
    <div className="flex flex-col lg:w-[530px] lg:h-[300px] md:w-[487px] md:h-[299px] h-[254]">
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
