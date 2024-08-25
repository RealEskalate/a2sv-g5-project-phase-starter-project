import React from "react";
import { GrSend } from "react-icons/gr";
export default function QuickTransfer() {
  return (
    <div className="relative flex-col  p-2  min-h-48 w-96 rounded-xl bg-white  text-white shadow-2xl ">
      <div className="text-black flex items-ceneter justify-center gap-5 overflow-auto mb-3">
        <div>
          <div className="avatar">
            <div className="w-20 h-20 rounded-full">
              <img src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
            </div>
          </div>
          <div className="flex justify-center">Work Man</div>
          <div className="flex justify-center font-thin">Designer</div>
        </div>
        <div>
          <div className="avatar">
            <div className="w-20 h-20 rounded-full">
              <img src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
            </div>
          </div>
          <div className="flex justify-center">Work Man</div>
          <div className="flex justify-center font-thin">Designer</div>
        </div>
        <div>
          <div className="avatar">
            <div className="w-20 h-20 rounded-full">
              <img src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
            </div>
          </div>
          <div className="flex justify-center">Work Man</div>
          <div className="flex justify-center font-thin">Designer</div>
        </div>
      </div>
      <div className="flex justify-between gap-3 items-center text-black">
        <p className="font-thin  ml-5">Write Amount</p>
        <div className="flex w-2/3 rounded-full    h-10  font-[sans-serif] text-black">
          <input
            placeholder="50.00..."
            className="w-full outline-none bg-slate-100 text-sm px-5 py-3"
          />
          <button
            type="button"
            className="flex items-center rounded-full w-15 text-white justify-center px-6 gap-3 bg-indigo-800"
          >
            <span>Send</span>
            <GrSend className="text-white text-xl" />
          </button>
        </div>
      </div>
    </div>
  );
}
