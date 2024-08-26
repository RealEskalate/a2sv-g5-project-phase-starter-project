import {
  faCheck,
  faDotCircle,
  faPen,
  faTimesSquare,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useState } from "react";

export default function ProgressComp({ currentStep }: { currentStep: any }) {
  const totalSteps = 3;
  const menuStep = [
    "Basic Information",
    "Address Information",
    "Personal Information",
  ];

  // Helper to determine if the circle should be active or not
  const isActive = (step: number) => step <= currentStep;

  return (
    <div className="progress-box flex flex-col  items-start justify-center gap-12">
      <div className="vert-line h-[208px] z-0 w-[2px] bg-gradient-to-b from-white to-blue-800 rounded-3xl absolute left-[60px]"></div>
      {[...Array(totalSteps)].map((_, index) => {
        const step = index + 2;
        return (
          <button
            key={step}
            disabled
            className={`circle z-20 font-semibold h-10 w-10 text-lg flex items-center justify-center rounded-full ${
              isActive(step)
                ? "bg-blue-800 text-white border-gray-50"
                : "bg-white text-blue-950 border-white"
            } border-2 border-solid`}
            // onClick={() => setCurrentStep(step)}
          >
            <FontAwesomeIcon
              icon={isActive(step) ? faCheck : faDotCircle}
              className={
                isActive(step) ? "text-white font-semibold" : "text-blue-900"
              }
            />
          </button>
        );
      })}
    </div>
  );
}
