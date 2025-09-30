# Ministry of Truth

*Full-stack news aggregation app with AI-powered satirical transformations using Go backend and OpenAI integration*

## Overview

Ministry of Truth is a sophisticated web application that demonstrates the power and dangers of misinformation through satirical Orwellian commentary. The application fetches real news from NewsAPI and uses OpenAI's GPT to transform headlines into dystopian propaganda in the style of George Orwell's 1984.

**Educational Purpose:** This project is designed for educational and satirical purposes to highlight the importance of critical thinking about news consumption and AI-generated content.

## Live Application

- **Frontend**: [https://ministry-of-truth.netlify.app](https://ministry-of-truth.netlify.app)
- **Backend API**: [https://ministry-of-truth.onrender.com](https://ministry-of-truth.onrender.com)

## Features

- **Real-time News Aggregation** - Fetches breaking news from NewsAPI
- **Category Filtering** - Technology, Business, Sports, Entertainment, Health, Science
- **AI Content Transformation** - Uses OpenAI GPT-3.5-turbo for satirical content generation
- **Dual Reality Toggle** - Switch between "Real News" and "Ministry Version"
- **Search Functionality** - Search news articles by keywords
- **Cost Controls** - Built-in usage limits to manage API costs
- **Professional Security** - Environment variable management for API keys

## Tech Stack

### Backend (Deployed on Render)
- **Go** with Gorilla Mux router
- **REST API** architecture with CORS middleware
- **Environment-based** configuration for security
- **Health check** endpoints for monitoring

### Frontend (Deployed on Netlify)
- **HTML5, CSS3** with modern responsive design
- **Vanilla JavaScript** with async/await patterns
- **Professional dark theme** with glassmorphism effects
- **Google Fonts** (Orbitron & Exo 2) typography

### External APIs
- **NewsAPI** for real-time news data from multiple sources
- **OpenAI GPT-3.5-turbo** for content transformation

## Architecture

This application uses a **distributed deployment strategy** across two specialized platforms:

```
Frontend (Netlify)     ←→     Backend (Render)     ←→     External APIs
Static Site Hosting           Go Server                  NewsAPI + OpenAI
- HTML/CSS/JS                 - REST API endpoints       - Real-time data
- Global CDN                  - Business logic           - AI processing
- Auto-deploy                 - Secure API keys          - Content transformation
```

### Why Two Platforms?

**Netlify** (Frontend):
- Optimized for static file delivery via global CDN
- Automatic deployments from GitHub
- Superior performance for HTML/CSS/JS assets

**Render** (Backend):
- Designed for server-side applications
- Native Go runtime support
- Secure environment variable management
- Always-on server processing

This separation follows industry best practices for scalability, security, and performance.

## Getting Started

### Prerequisites

- Go 1.19 or higher
- NewsAPI account and API key
- OpenAI account and API key

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/MitchellLeybaCale/ministry-of-truth.git
   cd ministry-of-truth
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   
   Copy `.env.example` to `.env` and fill in your API keys:
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` with your actual API keys:
   ```env
   NEWS_API_KEY=your_newsapi_key_here
   OPENAI_API_KEY=your_openai_key_here
   PORT=8080
   ENVIRONMENT=development
   ```

4. **Get API Keys**
   
   **NewsAPI:**
   - Visit [NewsAPI.org](https://newsapi.org)
   - Sign up for a free account
   - Copy your API key
   
   **OpenAI:**
   - Visit [OpenAI Platform](https://platform.openai.com)
   - Create an account and add billing
   - Generate an API key

5. **Run the application**
   ```bash
   go run main.go
   ```

6. **Open your browser**
   
   Navigate to `http://localhost:8080`

## Deployment Journey

This project demonstrates a complete deployment workflow from development to production:

### Backend Deployment (Render)
1. **Connected GitHub repository** to Render
2. **Configured build settings**:
   - Build Command: `go build -o app main.go`
   - Start Command: `./app`
   - Environment variables for API keys
3. **Automatic deployments** on git push to main branch
4. **Health monitoring** via `/health` endpoint

### Frontend Deployment (Netlify)
1. **Connected GitHub repository** to Netlify
2. **Configured build settings**:
   - Publish directory: `/` (root)
   - Auto-deploy from main branch
3. **Updated API endpoints** to point to Render backend
4. **Resolved CORS and routing** issues between platforms

### Development Challenges Overcome
- **Git workflow issues** with divergent branches
- **API endpoint mismatches** between frontend and backend
- **CORS configuration** for cross-platform communication
- **Layout and spacing** optimization for professional UI
- **Environment variable** management across platforms

## Project Structure

```
ministry-of-truth/
├── main.go              # Main Go backend server
├── public/              # Frontend files (deployed to Netlify)
│   └── index.html       # Main frontend application
├── go.mod              # Go module dependencies
├── go.sum              # Go module checksums
├── .env.example        # Environment variables template
├── .gitignore          # Git ignore rules
└── README.md           # Project documentation
```

## API Endpoints

- `GET /api/news/headlines` - Get top headlines
- `GET /api/news/headlines?category=technology` - Get categorized news
- `GET /api/news/search?q=keyword` - Search news articles
- `POST /api/transform` - Transform news content (OpenAI)
- `GET /health` - Health check endpoint

## Security Features

- **Environment Variables** - All API keys stored securely
- **API Key Masking** - Logs show `[REDACTED]` instead of actual keys
- **Git Protection** - `.gitignore` prevents accidental key commits
- **CORS Configuration** - Proper cross-origin request handling

## Cost Management

- **Daily Usage Limits** - Built-in limits to control OpenAI costs
- **Fallback Content** - Sample transformations when limits are reached
- **Usage Tracking** - Monitor API usage in development
- **Cost Estimation** - Approximately $0.09 maximum per day in development

## Design Philosophy

**Visual Theme:** "Classy Dark Sophistication"
- Deep blacks (#0a0a0a) and charcoals (#1a1a1a)
- Orbitron and Exo 2 font families with clean typography
- Smooth transitions and hover effects
- Grid-based responsive design with proper centering
- Glassmorphism effects for modern appeal

## Development Journey

This project represents a complete learning journey from API fundamentals to production-ready full-stack development:

### Phase 1: API Fundamentals (July 2025)
- Mastered API basics with Postman and JSONPlaceholder
- Built first HTML/JavaScript API integration
- Learned HTTP methods, authentication, and error handling

### Phase 2: Real-World Integration
- Connected to OpenWeatherMap API with authentication
- Integrated NewsAPI for live news aggregation  
- Added OpenAI for AI-powered content transformation

### Phase 3: Backend Development
- Evolved from single HTML file to Go backend architecture
- Implemented RESTful API with Gorilla Mux routing
- Added environment-based configuration and security

### Phase 4: Production Deployment
- **Overcame Render deployment challenges** with proper Go configuration
- **Resolved Netlify integration issues** with API endpoint mapping
- **Fixed Git workflow problems** with branch management
- **Optimized frontend layout** with proper CSS centering and spacing

### Phase 5: Cross-Platform Integration
- **Debugged CORS issues** between Netlify frontend and Render backend
- **Resolved API endpoint mismatches** between client and server
- **Implemented proper error handling** for network connectivity
- **Achieved full production deployment** across two platforms

## Key Technical Lessons

### API Integration
- Always test APIs in Postman before coding implementation
- Implement comprehensive error handling from the beginning
- Plan for rate limits and cost management early in development
- Use environment variables for all sensitive configuration

### Frontend-Backend Communication
- CORS configuration is critical for cross-platform deployment
- API endpoint consistency between client and server is essential
- Proper error handling provides better user experience
- Network debugging tools are invaluable for troubleshooting

### Deployment & DevOps
- Different platforms have specific requirements and limitations
- Git workflow management becomes critical in production
- Environment variable handling varies between development and production
- Testing deployment early prevents major integration issues

### UI/UX Development
- CSS centering requires understanding of flexbox and grid systems
- Responsive design needs mobile-first approach
- User feedback is crucial for asynchronous operations
- Professional polish makes a significant difference in perceived quality

## Alternative Deployment Options

While this project uses Netlify + Render, other viable options include:

**Single Platform Options:**
- **Vercel** - Excellent for Next.js and serverless functions
- **Railway** - Great for full-stack applications (our original choice)
- **Fly.io** - Good performance with global edge deployment

**Enterprise Options:**
- **AWS** (S3 + Lambda + API Gateway)
- **Google Cloud Platform**
- **Microsoft Azure**

The current setup optimizes for cost, performance, and learning experience.

## Contributing

This is a portfolio project, but feedback and suggestions are welcome! Please feel free to:
- Open issues for bugs or feature requests
- Submit pull requests for improvements
- Share your thoughts on the educational approach

## License

This project is open source and available under the [MIT License](LICENSE).

## Disclaimer

This application is created for educational and satirical purposes only. The "Ministry Version" content is clearly marked as AI-generated satire inspired by George Orwell's 1984. Users should always verify news from original sources and think critically about AI-generated content.

## Acknowledgments

- **George Orwell** - For the timeless warnings in "1984"
- **NewsAPI** - For providing excellent news aggregation services
- **OpenAI** - For the powerful GPT models enabling creative content transformation
- **Go Community** - For the excellent tools and documentation
- **Render & Netlify** - For reliable and developer-friendly deployment platforms

---

**Built by Mitchell Leyba Cale**

*"In a time of deceit telling the truth is a revolutionary act." - Often attributed to George Orwell*
