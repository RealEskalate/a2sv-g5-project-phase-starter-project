import React from "react";
import ServiceSkeletoncard from "./serviceSkeletoncard";

const ServiceSkeletonList = () => {
  return (
    <div
      className="flex overflow-x-auto justify-around overflow-clip whitespace-nowrap w-full"
      style={{
        scrollbarWidth: "none",
        msOverflowStyle: "none",
      }}
    >
      <ServiceSkeletoncard />
      <ServiceSkeletoncard />
      <ServiceSkeletoncard />
    </div>
  );
};

export default ServiceSkeletonList;
