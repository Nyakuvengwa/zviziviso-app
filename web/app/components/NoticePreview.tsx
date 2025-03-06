"use client";

import TraditionalTemplate from "./templates/TraditionalTemplate";
import CommunityTemplate from "./templates/CommunityTemplate";
import ModernTemplate from "./templates/ModernTemplate";
import { DeathNotice, TemplateType } from "../types";

interface NoticePreviewProps {
  notice: DeathNotice;
  template: TemplateType;
}

export default function NoticePreview({ notice, template }: NoticePreviewProps) {
  switch (template) {
    case "traditional":
      return <TraditionalTemplate notice={notice} />;
    case "community":
      return <CommunityTemplate notice={notice} />;
    case "modern":
      return <ModernTemplate notice={notice} />;
    default:
      return null;
  }
}