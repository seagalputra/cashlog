import React from "react";
import classNames from "classnames";
import { GenericComponentProps } from "types/ComponentProps";

export default function Header({
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

Header.Item = function HeaderItem({
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
