"use client";

import { Button } from "@/components/ui/button";
import { ToastAction } from "@/components/ui/toast";
import { useToast } from "@/components/ui/use-toast";

export function Toast() {
  const { toast } = useToast();

  return (
    <Button
      variant="outline"
      onClick={() => {
        console.log("toooooooooo");
        toast({
          title: "I got you",
          description: "Friday, February 10, 2023 at 5:57 PM",
          action: (
            <ToastAction altText="Goto schedule to undo">Undo</ToastAction>
          ),
        });
      }}
    >
      Add to calendar
    </Button>
  );
}
