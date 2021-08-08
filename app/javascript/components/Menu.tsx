import React from "react";
import classNames from "classnames";
import { GenericComponentProps } from "types/ComponentProps";

export default function Menu({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames("grid grid-cols-4 mt-6 bg-white border", className)}
      {...restProps}
    >
      {children}
    </div>
  );
}

Menu.Item = function MenuItem({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames(
        "inline-flex flex-col items-center text-center p-4 relative py-4 px-2 cursor-pointer",
        className
      )}
      {...restProps}
    >
      {children}
    </div>
  );
};

Menu.Link = function MenuLink({
  className,
  children,
  to,
  ...restProps
}: GenericComponentProps & { to: string }) {
  return (
    <a
      href={to}
      className={classNames(
        "text-gray-900 font-semibold text-xs mt-2",
        className
      )}
      {...restProps}
    >
      {children}
    </a>
  );
};
