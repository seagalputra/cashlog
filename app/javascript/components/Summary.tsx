import React from "react";
import classNames from "classnames";
import { GenericComponentProps } from "types/ComponentProps";

export default function Summary({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames("w-full p-4 bg-white border", className)}
      {...restProps}
    >
      {children}
    </div>
  );
}

Summary.Title = function SummaryTitle({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <p
      className={classNames("font-light text-gray-400", className)}
      {...restProps}
    >
      {children}
    </p>
  );
};

Summary.Item = function SummaryItem({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames("flex flex-col gap-2", classNames)}
      {...restProps}
    >
      {children}
    </div>
  );
};
