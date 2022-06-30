import { ReactNode } from "react";
import { Helmet } from "react-helmet-async";

interface Props {
  title: string;
  baseTitle?: string;
  children?: ReactNode | ReactNode[];
}

const BASE_TITLE = "Notes";

export function Head({ title, baseTitle, children }: Props) {
  return (
    <Helmet>
      <title>
        {title} {baseTitle ? ` - ${baseTitle}` : `- ${BASE_TITLE}`}
      </title>
      {children}
    </Helmet>
  );
}
