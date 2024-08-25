import React from "react";

export default function CardSettingItem({ Icon, title, description }: any) {
  return (
    <div className="flex bg-white items-center w-full overflow-clip">
      <div className="w-11 h-11 mr-2 flex-shrink-0">
        <Icon />
      </div>
      <div className="">
        <p className="text-slate-950 text-14px lg:text-15px">{title}</p>
        <p className="text-blue-steel text-10px lg:text-13px truncate">
          {description}
        </p>
      </div>
    </div>
  );
}
