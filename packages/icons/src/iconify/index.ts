import { addCollection, createIconifyIcon } from '@vben-core/icons';

import emojioneMonotone from './emojione-monotone.json';
import gridicons from './gridicons.json';
import lineMd from './line-md.json';
import lucide from './lucide.json';
import materialSymbols from './material-symbols.json';
import skillIcons from './skill-icons.json';

addCollection(emojioneMonotone);
addCollection(skillIcons);
addCollection(lucide);
addCollection(gridicons);
addCollection(lineMd);
addCollection(materialSymbols);

export * from '@vben-core/icons';

export const MdiKeyboardEsc = createIconifyIcon('mdi:keyboard-esc');

export const MdiWechat = createIconifyIcon('mdi:wechat');

export const MdiGithub = createIconifyIcon('mdi:github');

export const MdiGoogle = createIconifyIcon('mdi:google');

export const MdiQqchat = createIconifyIcon('mdi:qqchat');
