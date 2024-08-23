import React from "react";
import { GrSend } from "react-icons/gr";
export default function QuickTransfer() {
  return (
    <div className="relative flex-col    h-48 w-96 rounded-xl bg-gradient-to-r  text-white shadow-2xl ">
      <div className="text-black flex items-ceneter justify-center gap-2 overflow-hidden">
        <div className="rounded-full bg-slate-300">P1</div>{" "}
        {/* this is just a place holder */}
        <div className="rounded-full bg-slate-300">P1</div>
        <div className="rounded-full bg-slate-300">P1</div>
        <div className="rounded-full bg-slate-300">P1</div>
        <div className="rounded-full bg-slate-300">P1</div>
        <div className="rounded-full bg-slate-300">P1</div>
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
