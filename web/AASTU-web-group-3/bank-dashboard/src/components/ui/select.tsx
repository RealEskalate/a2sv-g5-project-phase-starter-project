import * as React from "react";

interface SelectProps {
  value: string;
  onValueChange: (value: string) => void;
  children: React.ReactNode;
}

interface SelectTriggerProps {
  className?: string;
  children: React.ReactNode;
  ariaLabel?: string;
}

interface SelectContentProps {
  children: React.ReactNode;
  align?: string;
  className?: string;
}

export function Select({ value, onValueChange, children }: SelectProps) {
  return (
    <select
      value={value}
      onChange={(e) => onValueChange(e.target.value)}
      className="border rounded px-2 py-1"
    >
      {children}
    </select>
  );
}

export function SelectTrigger({
  children,
  className,
  ariaLabel,
}: SelectTriggerProps) {
  return (
    <div className={className} aria-label={ariaLabel}>
      {children}
    </div>
  );
}

export function SelectValue({ placeholder }: { placeholder: string }) {
  return <div className="select-value">{placeholder}</div>;
}

export function SelectContent({
  children,
  align,
  className,
}: SelectContentProps) {
  return (
    <div className={`${className} ${align && `align-${align}`}`}>
      {children}
    </div>
  );
}

export function SelectItem({
  value,
  children,
  className,
}: {
  value: string;
  children: React.ReactNode;
  className?: string;
}) {
  return (
    <option value={value} className={className}>
      {children}
    </option>
  );
}
