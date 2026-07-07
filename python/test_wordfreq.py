"""Baseline sanity tests for wordfreq.

These exercise behavior that is orthogonal to the three known bugs (they use
lowercase, unpunctuated input and never request more rows than exist), so they
stay green on the current code and give slopstop a clean regression baseline to
start from.
"""
from wordfreq import top_words, word_frequencies


def test_counts_simple_words():
    assert word_frequencies("cat cat dog") == {"cat": 2, "dog": 1}


def test_empty_input_has_no_words():
    assert word_frequencies("") == {}


def test_top_words_orders_by_count_then_alphabetically():
    counts = {"cat": 2, "dog": 1, "ant": 1}
    assert top_words(counts, 10) == [("cat", 2), ("ant", 1), ("dog", 1)]
