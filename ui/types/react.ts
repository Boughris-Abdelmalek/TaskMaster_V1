import { FC, PropsWithChildren } from "react";

export type FCWithChildren<P = Record<string, unknown>> = FC<PropsWithChildren<P>>;
