document.addEventListener('DOMContentLoaded', () => {
    const tagPills = document.querySelectorAll('.tag-filter-pill');
    const articleCards = document.querySelectorAll('.article-card');

    function filterArticles(selectedTag) {
        tagPills.forEach(pill => {
            if (pill.dataset.tag === selectedTag) {
                pill.classList.add('active');
            } else {
                pill.classList.remove('active');
            }
        });

        articleCards.forEach(card => {
            const tags = card.dataset.tags.split(' ').filter(Boolean);
            if (selectedTag === 'All' || tags.includes(selectedTag)) {
                card.style.display = '';
            } else {
                card.style.display = 'none';
            }
        });
    }

    tagPills.forEach(pill => {
        pill.addEventListener('click', () => {
            const selectedTag = pill.dataset.tag;
            filterArticles(selectedTag);
            if (selectedTag === 'All') {
                window.location.hash = '';
            } else {
                window.location.hash = selectedTag;
            }
        });
    });

    // Initial filter based on URL hash
    let initialHashTag = window.location.hash.substring(1);
    if (initialHashTag === '') {
        initialHashTag = 'All';
    }

    if (initialHashTag) {
        const foundPill = Array.from(tagPills).find(pill => pill.dataset.tag === initialHashTag);
        if (foundPill) {
            filterArticles(initialHashTag);
        } else {
            filterArticles('All');
        }
    } else {
        filterArticles('All');
    }
});
