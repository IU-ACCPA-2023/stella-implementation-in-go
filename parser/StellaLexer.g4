
lexer grammar StellaLexer;
// Predefined regular expressions in BNFC
fragment LETTER: CAPITAL | SMALL;
fragment CAPITAL: [A-Z\u00C0-\u00D6\u00D8-\u00DE];
fragment SMALL: [a-z\u00DF-\u00F6\u00F8-\u00FF];
fragment DIGIT: [0-9];
Surrogate_id_SYMB_0: ',';
Surrogate_id_SYMB_1: ';';
Surrogate_id_SYMB_2: '(';
Surrogate_id_SYMB_3: ')';
Surrogate_id_SYMB_4: '{';
Surrogate_id_SYMB_5: '}';
Surrogate_id_SYMB_6: '=';
Surrogate_id_SYMB_7: ':';
Surrogate_id_SYMB_8: '->';
Surrogate_id_SYMB_9: '=>';
Surrogate_id_SYMB_10: '<|';
Surrogate_id_SYMB_11: '|>';
Surrogate_id_SYMB_12: '[';
Surrogate_id_SYMB_13: ']';
Surrogate_id_SYMB_14: '<';
Surrogate_id_SYMB_15: '<=';
Surrogate_id_SYMB_16: '>';
Surrogate_id_SYMB_17: '>=';
Surrogate_id_SYMB_18: '==';
Surrogate_id_SYMB_19: '!=';
Surrogate_id_SYMB_20: '+';
Surrogate_id_SYMB_21: '*';
Surrogate_id_SYMB_22: 'List::head';
Surrogate_id_SYMB_23: 'List::isempty';
Surrogate_id_SYMB_24: 'List::tail';
Surrogate_id_SYMB_25: 'Nat::pred';
Surrogate_id_SYMB_26: 'Nat::iszero';
Surrogate_id_SYMB_27: 'Nat::rec';
Surrogate_id_SYMB_28: '.';
Surrogate_id_SYMB_29: 'Bool';
Surrogate_id_SYMB_30: 'Nat';
Surrogate_id_SYMB_31: 'Unit';
Surrogate_id_SYMB_32: 'and';
Surrogate_id_SYMB_33: 'as';
Surrogate_id_SYMB_34: 'cons';
Surrogate_id_SYMB_35: 'core';
Surrogate_id_SYMB_36: 'else';
Surrogate_id_SYMB_37: 'extend';
Surrogate_id_SYMB_38: 'false';
Surrogate_id_SYMB_39: 'fix';
Surrogate_id_SYMB_40: 'fn';
Surrogate_id_SYMB_41: 'fold';
Surrogate_id_SYMB_42: 'if';
Surrogate_id_SYMB_43: 'in';
Surrogate_id_SYMB_44: 'inline';
Surrogate_id_SYMB_45: 'language';
Surrogate_id_SYMB_46: 'let';
Surrogate_id_SYMB_47: 'match';
Surrogate_id_SYMB_48: 'not';
Surrogate_id_SYMB_49: 'or';
Surrogate_id_SYMB_50: 'record';
Surrogate_id_SYMB_51: 'return';
Surrogate_id_SYMB_52: 'succ';
Surrogate_id_SYMB_53: 'then';
Surrogate_id_SYMB_54: 'throws';
Surrogate_id_SYMB_55: 'true';
Surrogate_id_SYMB_56: 'type';
Surrogate_id_SYMB_57: 'unfold';
Surrogate_id_SYMB_58: 'variant';
Surrogate_id_SYMB_59: 'with';
Surrogate_id_SYMB_60: 'µ';
COMMENT_antlr_builtin: ('//' ~[\r\n]* (('\r'? '\n') | EOF)) -> skip;

StellaIdent: ('_' | LETTER) ([!\-:?_] | (DIGIT | LETTER))*;
ExtensionName: '#' ([\-_] | (DIGIT | LETTER))+;

//Integer predefined token type
INTEGER: DIGIT+;

// Whitespace
WS: (' ' | '\r' | '\t' | '\n' | '\f')+ -> skip;
// Escapable sequences
fragment Escapable: ('"' | '\\' | 'n' | 't' | 'r' | 'f');
ErrorToken: .;
