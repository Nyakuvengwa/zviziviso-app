export interface DeathNotice {
  name: string;
  age: number;
  dateOfDeath: string;
  locationOfDeath: string;
  familyInfo: string;
  funeralArrangements: {
    wake: string;
    service: string;
    burial: string;
  };
  religion: string;
  tributes: string;
  imageUrl?: string;
}

export interface PublishedNotice {
  id: string;
  notice: DeathNotice;
  template: TemplateType;
  publishedAt: string;
}

export type TemplateType = 'traditional' | 'community' | 'modern';