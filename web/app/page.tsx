"use client";

import { useState, useEffect } from "react";
import NoticeForm from "./components/NoticeForm";
import NoticeList from "./components/NoticeList";
import NoticePreview from "./components/NoticePreview";
import { DeathNotice, PublishedNotice, TemplateType } from "./types";

export default function Home() {
  const [notices, setNotices] = useState<PublishedNotice[]>([]);
  const [draftNotice, setDraftNotice] = useState<{
    notice: DeathNotice;
    template: TemplateType;
  } | null>(null);
  const [selectedNotice, setSelectedNotice] = useState<PublishedNotice | null>(null);

  useEffect(() => {
    const savedNotices = localStorage.getItem('deathNotices');
    if (savedNotices) {
      setNotices(JSON.parse(savedNotices));
    }
  }, []);

  const handleSubmit = (notice: DeathNotice, template: TemplateType) => {
    setDraftNotice({ notice, template });
  };

  const handlePublish = () => {
    if (!draftNotice) return;

    const newNotice: PublishedNotice = {
      id: crypto.randomUUID(),
      notice: draftNotice.notice,
      template: draftNotice.template,
      publishedAt: new Date().toISOString(),
    };

    const updatedNotices = [newNotice, ...notices];
    setNotices(updatedNotices);
    localStorage.setItem('deathNotices', JSON.stringify(updatedNotices));
    setDraftNotice(null);
  };

  const handleNoticeClick = (notice: PublishedNotice) => {
    setSelectedNotice(notice);
    setDraftNotice(null);
  };

  return (
    <main className="min-h-screen bg-[#f8f5f1] py-12">
      <div className="container mx-auto px-4">
        <header className="text-center max-w-4xl mx-auto mb-12">
          <h1 className="text-4xl font-serif text-gray-900 mb-2">Zimbabwean Death Notices</h1>
          <p className="text-gray-600">Create and share respectful and culturally appropriate death notices</p>
        </header>

        {!draftNotice && !selectedNotice && (
          <div className="space-y-12">
            <div className="max-w-4xl mx-auto">
              <NoticeForm onSubmit={handleSubmit} />
            </div>
            {notices.length > 0 && (
              <div className="max-w-6xl mx-auto">
                <NoticeList notices={notices} onNoticeClick={handleNoticeClick} />
              </div>
            )}
          </div>
        )}

        {draftNotice && (
          <div className="max-w-4xl mx-auto space-y-6">
            <div className="flex items-center justify-between">
              <button
                onClick={() => setDraftNotice(null)}
                className="text-gray-600 hover:text-gray-900 transition-colors"
              >
                ← Back to Form
              </button>
              <button
                onClick={handlePublish}
                className="bg-green-600 text-white px-6 py-2 rounded-lg hover:bg-green-700 transition-colors"
              >
                Publish Notice
              </button>
            </div>
            <NoticePreview notice={draftNotice.notice} template={draftNotice.template} />
          </div>
        )}

        {selectedNotice && (
          <div className="max-w-4xl mx-auto space-y-6">
            <button
              onClick={() => setSelectedNotice(null)}
              className="text-gray-600 hover:text-gray-900 transition-colors"
            >
              ← Back to Notices
            </button>
            <NoticePreview notice={selectedNotice.notice} template={selectedNotice.template} />
          </div>
        )}
      </div>
    </main>
  );
}