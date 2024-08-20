import { useState } from "react";

export default function ProgressComp() {
  const [currentStep, setCurrentStep] = useState(1);
  const totalSteps = 3;

  // Helper to determine if the circle should be active or not
  const isActive = (step: number) => step <= currentStep;

  return (
    <div className="progress-box w-full flex flex-col  items-start justify-center gap-12 px-12 py-4">
      {[...Array(totalSteps)].map((_, index) => {
        const step = index + 1;
        return (
          <div
            key={step}
            className={`circle h-10 w-10 text-lg flex items-center justify-center rounded-full ${
              isActive(step)
                ? "bg-blue-600 text-white border-blue-600"
                : "bg-blue-50 text-blue-500 border-blue-600"
            } border-2 border-solid cursor-pointer`}
            onClick={() => setCurrentStep(step)}
          >
            {step}
          </div>
        );
      })}
    </div>
  );
}
