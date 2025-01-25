type Theme = 'dark' | 'light' | 'system';

let theme = $state<Theme>('system');

export function setTheme(newTheme: Theme) {
	theme = newTheme;
	updateTheme();
}

export function getTheme(): Theme {
	return theme;
}

function updateTheme() {
	const root = document.documentElement;
	const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
	const activeTheme = theme === 'system' ? systemTheme : theme;

	root.classList.remove('light', 'dark');
	root.classList.add(activeTheme);
}
