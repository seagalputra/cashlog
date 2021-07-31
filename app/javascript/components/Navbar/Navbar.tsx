import React from "react";
import classNames from "classnames";

type NavbarProps = {
  className?: any;
  children: any;
};

export default function Navbar({
  className,
  children,
  ...restProps
}: NavbarProps) {
  return (
    <nav
      className={classNames(
        "fixed inset-x-0 top-0 z-50 flex items-center justify-between w-full p-3 bg-gray-800",
        className
      )}
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
}: NavbarProps) {
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
}: NavbarProps) {
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
