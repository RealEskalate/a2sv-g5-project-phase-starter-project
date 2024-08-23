import { useState } from "react";

export default function ProgressComp({ currentStep }: { currentStep: any }) {
  const totalSteps = 3;
  const menuStep = ["Basic Information","Address Information","Personal Information"]

  // Helper to determine if the circle should be active or not
  const isActive = (step: number) => step <= currentStep;

  return (
    <div className="progress-box w-full flex flex-col items-start justify-center  px-12">
      {[...Array(totalSteps)].map((_, index) => {
        const step = index + 1;
        return (
          <div key={step} className="flex flex-col items-center">
            <div className="w-[210px] flex justify-between items-center">
              <h1 
                className={`text-base ${
                  isActive(step)
                  ? "text-blue-700 glow-effect"
                    : "text-blue-300 "
                } `}>
                  {menuStep[index]}
                </h1>
              <div
                className={`circle h-10 w-10 text-lg flex items-center justify-center rounded-full ${
                  isActive(step)
                    ? "bg-blue-600 text-white border-blue-600 glow-effect"
                    : "bg-blue-50 text-blue-500 border-blue-600"
                } border-2 border-solid cursor-pointer`}
              >
                {step}
              </div>
            </div>
            {step < totalSteps && (
              <div className={`h-8 w-0.5 my-2 self-end mr-5 ${
                isActive(step+1)
                ? "bg-blue-600"
                : "bg-blue-300"
              } `}></div>
            )}
          </div>
        );
      })}
    </div>
  );
}
