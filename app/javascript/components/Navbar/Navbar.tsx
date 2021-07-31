import React from "react";
import classNames from "classnames";
import { GenericComponentProps } from "types/ComponentProps";
export default function Navbar({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <nav
      className={classNames("fixed inset-x-0 top-0 z-50 w-full p-3", className)}
      {...restProps}
    >
      {children}
    </nav>
  );
}

Navbar.Title = function NavbarTitle({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames("flex items-center space-x-6", className)}
      {...restProps}
    >
      {children}
    </div>
  );
};

Navbar.Item = function NavbarItem({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames("flex items-center gap-2", className)}
      {...restProps}
    >
      <div className="ml-3 relative">
        <div>{children}</div>
      </div>
    </div>
  );
};
