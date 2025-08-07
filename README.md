# Ministry of Truth

> **Educational satirical project demonstrating AI-powered news transformation**

A full-stack web application that transforms real news headlines into dystopian propaganda in the style of George Orwell's "1984" using AI, built to explore and educate about misinformation techniques.

## [Live Demo](https://ministry-of-truth-production.up.railway.app)

![Ministry of Truth Screenshot](https://via.placeholder.com/1200x600/0a0a0a/58a6ff?text=Ministry+of+Truth+Screenshot)

## Project Overview

**Ministry of Truth** is a sophisticated news aggregation platform that demonstrates the power and dangers of information manipulation through satirical Orwellian commentary. Users can:

- Browse real news from multiple categories (Technology, Business, Sports, etc.)
- Search for specific topics and articles
- Toggle between "Real News" and "Ministry Version" to see AI-transformed propaganda
- Experience a professional, Google-style dark UI with smooth animations

**Important**: All transformed content is clearly marked as fictional and satirical for educational purposes.

## Technical Stack

### Backend
- **Go** - High-performance server with Gorilla Mux routing
- **Railway** - Cloud deployment platform
- **RESTful APIs** - Clean API architecture

### Frontend
- **HTML5/CSS3** - Modern responsive design with CSS Grid/Flexbox
- **Vanilla JavaScript** - Professional architecture with async/await
- **Google Fonts** - Orbitron & Exo 2 typography

### External APIs
- **NewsAPI** - Real-time news aggregation from multiple sources
- **OpenAI GPT-3.5-turbo** - AI-powered content transformation

### Key Features
- **Smart Caching** - 24-hour transformation cache to optimize API usage
- **Cost Controls** - Built-in limits to prevent runaway API costs
- **CORS Configuration** - Proper cross-origin resource sharing
- **Responsive Design** - Mobile-first approach with breakpoints
- **Error Handling** - Comprehensive error states and user feedback

## Screenshots

### Real News Mode
*Clean, professional news browsing with category filtering*

### Ministry Version Mode
*AI-transformed headlines with clear satirical labeling*

### Mobile Experience
*Fully responsive design optimized for all devices*

## Learning Journey

This project represents a complete journey from **zero API knowledge to full-stack deployment**:

### Phase 1: API Fundamentals (July 2025)
- Started with Postman and JSONPlaceholder for practice
- Learned HTTP methods, headers, and authentication
- Built first working HTML/JavaScript API application

### Phase 2: Real-World Integration
- Connected to OpenWeatherMap API with authentication
- Integrated NewsAPI for live news aggregation
- Added OpenAI for AI-powered content transformation

### Phase 3: Production Architecture
- Evolved from single HTML file to professional frontend/backend separation
- Implemented Go server with Gorilla Mux routing
- Added smart caching and cost control mechanisms

### Phase 4: Deployment & DevOps
- Overcame deployment challenges on multiple platforms
- Configured environment variables and CORS properly
- Achieved successful Railway deployment with custom domain

## Architecture

```
Frontend (HTML/CSS/JS)
    ↓ API Calls
Go Backend Server
    ↓ External API Integration
NewsAPI + OpenAI APIs
    ↓ Deployment
Railway Cloud Platform
```

### API Endpoints

```
GET  /api/health                    # Health check
GET  /api/news/headlines           # Top headlines
GET  /api/news/headlines?category  # Category filtering  
GET  /api/news/search?q=query      # Search functionality
POST /api/transform                # AI transformation
```

## Getting Started

### Prerequisites
- Go 1.21+
- News API key from [newsapi.org](https://newsapi.org)
- OpenAI API key from [openai.com](https://platform.openai.com)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/MitchellLeybaCale/ministry-of-truth.git
   cd ministry-of-truth
   ```

2. **Set environment variables**
   ```bash
   export NEWS_API_KEY="your_newsapi_key"
   export OPENAI_API_KEY="your_openai_key" 
   export PORT="8080"
   ```

3. **Install dependencies and run**
   ```bash
   go mod tidy
   go run main.go
   ```

4. **Visit the application**
   ```
   http://localhost:8080
   ```

### Deployment

The application is configured for easy Railway deployment:

1. Connect your GitHub repository to Railway
2. Set environment variables in Railway dashboard
3. Automatic deployment on git push

## Key Technical Decisions

### Why Go for the Backend?
- **Performance**: Fast compilation and execution
- **Simplicity**: Clean syntax and excellent standard library
- **Concurrency**: Built-in goroutines for handling multiple requests
- **Deployment**: Single binary deployment to Railway

### Why Vanilla JavaScript?
- **Performance**: No framework overhead
- **Learning**: Direct DOM manipulation and API interaction
- **Simplicity**: Easier to understand and debug
- **Portability**: Works everywhere without build steps

### Cost Management Strategy
- **Smart Caching**: 24-hour localStorage cache for transformations
- **Usage Limits**: 15 transformations per user per day
- **Fallback Content**: Graceful degradation when limits reached
- **Global Limits**: 50 transformations per day across all users

## Security & Privacy

- **API Key Protection**: All keys stored securely in environment variables
- **CORS Configuration**: Proper cross-origin resource sharing
- **Input Validation**: Server-side validation for all user inputs
- **Rate Limiting**: Built-in protection against abuse

## UI/UX Philosophy

**"Classy Dark Sophistication"**
- Professional dark theme inspired by modern developer tools
- Smooth animations and micro-interactions
- Google-style layout with clean typography
- Accessible color contrast and responsive design
- Clear visual hierarchy and information architecture

## Lessons Learned

### API Integration
- Always test APIs in Postman before coding
- Implement comprehensive error handling from day one
- Plan for rate limits and cost management early
- Use environment variables for all configuration

### Frontend Development  
- Vanilla JavaScript is powerful for many use cases
- Responsive design requires mobile-first thinking
- User feedback is crucial for async operations
- Caching strategies can dramatically improve UX

### Backend Development
- Go's simplicity makes it excellent for API servers
- CORS configuration is critical for frontend-backend separation
- Structured logging helps with debugging deployment issues
- Environment-based configuration enables smooth deployments

### Deployment & DevOps
- Platform-specific requirements vary significantly
- Git workflow and version control are essential
- Environment variable management is crucial
- Testing deployment early prevents last-minute issues

## Future Enhancements

- [ ] User authentication and personalized settings
- [ ] Multiple AI model support (GPT-4, Claude, etc.)
- [ ] Custom domain with HTTPS
- [ ] Advanced analytics and usage tracking  
- [ ] Social sharing capabilities
- [ ] Bookmarking and favorites system
- [ ] Docker containerization
- [ ] CI/CD pipeline with automated testing

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This is an educational project created to demonstrate web development skills and explore the topic of misinformation in a satirical context. All "Ministry Version" content is clearly marked as fictional and generated by AI for educational purposes. The project is not intended to spread misinformation or deceive users.

## About the Developer

Built by **Mitchell Leyba-Cale** as part of a learning journey from zero API knowledge to full-stack deployment. This project demonstrates:

- **Full-stack development** capabilities
- **API integration** expertise  
- **Modern JavaScript** proficiency
- **Go backend** development
- **Cloud deployment** experience
- **Professional development** practices

---

**If you found this project interesting or helpful, please consider giving it a star!**

**Questions or feedback?** Feel free to [open an issue](https://github.com/MitchellLeybaCale/ministry-of-truth/issues) or connect with me on LinkedIn.
