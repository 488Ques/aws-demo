async function deleteProduct(deleteEndpoint) {
    await fetch(deleteEndpoint, {
        method: 'DELETE',
    });
    location.reload();
}