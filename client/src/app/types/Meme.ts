export type MemeSize = 'small' | 'medium' | 'large';

export interface Meme {
  id: number;
  media_url: string;
  media_type: string;
  tags: string[];
}