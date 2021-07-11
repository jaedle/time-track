import { Selector } from 'testcafe';

fixture`Links`.page`http://localhost:4200/`;

test('shows hello world', async (t) => {
  await t.expect(Selector('body').innerText).contains('Hello World');
});
