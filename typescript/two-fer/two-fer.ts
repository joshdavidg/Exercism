/*************** TypeScript Only Implementation **************/
type EmptyOrNothing = "" | null | undefined;
type StringOrNothing = string | null | undefined;
type HasName<Name extends StringOrNothing> = Name extends EmptyOrNothing ? "One for you, one for me." : `One for ${Name}, one for me.`;
type CheckName = "Bob";
type StringWithName = HasName<CheckName>; // Hover over StringWithName to see the string with a name
type StringNoName = HasName<null>; // Hover over StringNoName to see the string without a name
/************* End TypeScript Only Implementation ************/


export function twoFer(name: string = "you"): string {
    return `One for ${name}, one for me.`;
}
