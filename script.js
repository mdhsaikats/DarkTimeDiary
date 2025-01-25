const photos = [
    { id: 1, src: "https://picsum.photos/300?random=1", title: "Photo 1", description: "A beautiful view of the mountains during sunset." },
    { id: 2, src: "https://picsum.photos/300?random=2", title: "Photo 2", description: "A serene beach with crystal-clear waters." },
    { id: 3, src: "https://picsum.photos/300?random=3", title: "Photo 3", description: "A bustling cityscape at night with bright lights." },
    { id: 4, src: "https://picsum.photos/300?random=4", title: "Photo 4", description: "A calm forest path surrounded by tall trees." },
    { id: 5, src: "https://picsum.photos/300?random=5", title: "Photo 5", description: "A vibrant field of flowers in full bloom." },
    { id: 6, src: "https://picsum.photos/300?random=6", title: "Photo 6", description: "A snowy landscape with a cozy cabin in the distance." }
  ];
  
  const gallery = document.getElementById('photo-gallery');
  
  function loadPhotos() {
    photos.forEach(photo => {
      const card = document.createElement('div');
      card.classList.add('card');
  
      card.innerHTML = `
        <img src="${photo.src}" alt="${photo.title}">
        <div class="card-content">
          <h2 class="card-title">${photo.title}</h2>
          <p class="card-description">${photo.description}</p>
        </div>
      `;
      gallery.appendChild(card);
    });
  }
  
  document.getElementById('load-more').addEventListener('click', () => {
    loadPhotos();
  });
  
  // Load initial photos
  loadPhotos();
  
  