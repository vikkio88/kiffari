import { expect, test } from 'vitest';
import { del, get, login, logout, post, put } from '../utils/api';
import { randomString } from '../utils/strings';

const rubbishTags = [];


test("Config", async () => {
    const resp = await get("config");
    expect(resp.ok).toBe(true);

    const data = await resp.json();
    expect(data.version).toBe("DEV");

});

test("login", async () => {
    const WRONG_PASSWORD = "wrongPassword";
    let resp = await post("login", { passKey: WRONG_PASSWORD });
    expect(resp.ok).toBe(false);
    expect(resp.status).toBe(401);

    const CORRECT_PASSWORD = "password";
    resp = await post("login", { passKey: CORRECT_PASSWORD });
    expect(resp.ok).toBe(true);

    const data = await resp.json();
    expect(data.token).toContain("ey");
});

test("notes / tags", async () => {
    logout();
    let resp = await get("notes");
    expect(resp.ok).toBe(false);
    expect(resp.status).toBe(401);
    await login();

    // GET ALL
    resp = await get("notes");
    expect(resp.ok).toBe(true);


    // CREATE
    const newNote = {
        title: "Some Title", body: "Some Body", tags: [
            { label: `TESTING_${randomString()}` }
        ]
    };
    resp = await post("notes", newNote);
    expect(resp.ok).toBe(true);
    let result = await resp.json();
    delete newNote.tags;
    expect(result).toMatchObject(newNote);

    // GET ONE
    const createdId = result.id;
    resp = await get(`notes/${createdId}`);
    expect(resp.ok).toBe(true);
    result = await resp.json();

    expect(result).toMatchObject(newNote);

    const fullNote = { ...result };

    const toUpdateNote = {
        id: fullNote.id,
        title: "Updated Title",
        body: fullNote.body,
        tags: [
            ...fullNote.tags,
            { label: `TESTING_${randomString()}` }
        ],
        // UPDATE NEEDS ARCHIVED TOO
        archived: false,
        due_date: null,
    };
    // UPDATE
    resp = await put(`notes/${fullNote.id}`, toUpdateNote);
    expect(resp.ok).toBe(true);
    resp = await get(`notes/${result.id}`);
    result = await resp.json();
    expect(result.title).toBe("Updated Title");
    expect(result.tags.length).toBe(2);

    // Archive
    resp = await get("notes?archived=true");
    expect(resp.ok).toBe(true);
    let archived = await resp.json();
    expect(archived.length).toBe(0);

    resp = await post(`notes/${fullNote.id}/archive`);
    expect(resp.ok).toBe(true);
    //Get only archived
    resp = await get("notes?archived=true");
    expect(resp.ok).toBe(true);
    archived = await resp.json();
    expect(archived.length).toBe(1);
    // un-archive
    resp = await del(`notes/${fullNote.id}/archive`);
    expect(resp.ok).toBe(true);
    //Get only archived
    resp = await get("notes?archived=true");
    expect(resp.ok).toBe(true);
    archived = await resp.json();
    expect(archived.length).toBe(0);

    // Pin/Unpin
    resp = await post(`notes/${fullNote.id}/pin`);
    expect(resp.ok).toBe(true);
    resp = await get(`notes/${result.id}`);
    result = await resp.json();
    expect(result.pinned).toBe(true);

    resp = await del(`notes/${fullNote.id}/pin`);
    expect(resp.ok).toBe(true);
    resp = await get(`notes/${result.id}`);
    result = await resp.json();
    expect(result.pinned).toBe(false);

    // Dash
    resp = await get('dash/notes');
    expect(resp.ok).toBe(true);
    result = await resp.json();
    expect(Array.isArray(result.latest)).toBe(true);
    expect(Array.isArray(result.pinned)).toBe(true);
    expect(Array.isArray(result.reminders)).toBe(true);


    // DELETE
    resp = await del(`notes/${fullNote.id}`);
    expect(resp.ok).toBe(true);

    // Get Again to check
    resp = await get(`notes/${fullNote.id}`);
    expect(resp.ok).not.toBe(true);
    expect(resp.status).toBe(404);

    // remove all testing tags
    resp = await get(`tags?q=TESTING`);
    expect(resp.ok).toBe(true);

    const tags = await resp.json();
    expect(tags.length).not.toBe(0);

    for (const t of tags) {
        resp = await del(`tags/${t.id}`);
        expect(resp.ok).toBe(true);
    }
});